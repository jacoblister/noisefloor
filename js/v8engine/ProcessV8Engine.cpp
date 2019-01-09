#include "ProcessV8Engine.hpp"

#include <iostream>

void ProcessV8Engine::compile(std::string source) {
    v8::Local<v8::String> source_func = v8::String::NewFromUtf8(isolate, source.c_str(), v8::NewStringType::kNormal).ToLocalChecked();
    v8::Local<v8::Script> script_func = v8::Script::Compile(this->context.Get(this->isolate), source_func).ToLocalChecked();
    script_func->Run(this->context.Get(this->isolate)).ToLocalChecked();
}

void ProcessV8Engine::compile_source(std::string filename) {
    std::ifstream t(filename.c_str());
    std::string source((std::istreambuf_iterator<char>(t)),
                    std::istreambuf_iterator<char>());
    v8::Local<v8::String> source_func = v8::String::NewFromUtf8(this->isolate, source.c_str(), v8::NewStringType::kNormal).ToLocalChecked();
    v8::Local<v8::Script> script_func = v8::Script::Compile(this->context.Get(this->isolate), source_func).ToLocalChecked();
    script_func->Run(this->context.Get(this->isolate)).ToLocalChecked();
}

// Extracts a C string from a V8 Utf8Value.
const char* ToCString(const v8::String::Utf8Value& value) {
  return *value ? *value : "<string conversion failed>";
}

// The callback that is invoked by v8 whenever the JavaScript 'console_log'
// function is called.  Prints its arguments on stdout separated by
// spaces and ending with a newline.
void ConsoleLog(const v8::FunctionCallbackInfo<v8::Value>& args) {
  bool first = true;
  for (int i = 0; i < args.Length(); i++) {
    v8::HandleScope handle_scope(args.GetIsolate());
    if (first) {
      first = false;
    } else {
      printf(" ");
    }
    v8::String::Utf8Value str(args.GetIsolate(), args[i]);
    const char* cstr = ToCString(str);
    printf("%s", cstr);
  }
  printf("\n");
  fflush(stdout);
}

result<bool> ProcessV8Engine::init(void) {
    printf("v8 init\n");

    return true;
}

result<bool> ProcessV8Engine::start(int sampling_rate, int samples_per_frame) {
    this->samples_per_frame = samples_per_frame;

    v8::V8::InitializeICUDefaultLocation("");
    v8::V8::InitializeExternalStartupData("");
    v8::Platform *platform = v8::platform::CreateDefaultPlatform();
    v8::V8::InitializePlatform(platform);
    v8::V8::Initialize();

    // Create a new Isolate and make it the current one.
    v8::Isolate::CreateParams create_params;
    create_params.array_buffer_allocator =
      v8::ArrayBuffer::Allocator::NewDefaultAllocator();
    this->isolate = v8::Isolate::New(create_params);
    v8::Isolate::Scope isolate_scope(this->isolate);

    // Create a stack-allocated handle scope.
    v8::HandleScope handle_scope(this->isolate);

    // Bind the global 'console_log' function to the C++ Print callback.
    v8::Local<v8::ObjectTemplate> global = v8::ObjectTemplate::New(isolate);
    global->Set(
      v8::String::NewFromUtf8(isolate, "console_log", v8::NewStringType::kNormal)
          .ToLocalChecked(),
      v8::FunctionTemplate::New(isolate, ConsoleLog));

    // Create a new context.
    v8::Local<v8::Context> local_context = v8::Context::New(this->isolate, NULL, global);
    this->context = v8::Eternal<v8::Context>(this->isolate, local_context);

    // Enter the context for compiling and running script.
    v8::Context::Scope context_scope(local_context);

    // Set console log function
    this->compile("console.log = function (message) { console_log(message) };");

    this->compile_source("../engine/engine.js");
    this->compile("function start(sampleRate) { noisefloorjs.start(sampleRate); }");
    this->compile("function process(samplesIn, samplesOut, midiIn, midiOut) { return noisefloorjs.process(samplesIn, samplesOut, midiIn, midiOut); }");
    // this->compile("function query(endpoint, request) { return engine.org.noisefloor.engine.query(endpoint, request); }");

    // this->compile_source("./gain.js");
    // this->compile("function process(samples) { return samples; }");

    v8::Local<v8::Value>    start_function_value = local_context->Global()->Get(v8::String::NewFromUtf8(isolate, "start"));
    v8::Local<v8::Function> start_function = v8::Local<v8::Function>::Cast(start_function_value);
    v8::Local<v8::Value>    args[] = { v8::Number::New(this->isolate, (double)sampling_rate) };
    start_function->Call(local_context->Global(), 1, args);

    v8::Local<v8::Value>    process_function_value = local_context->Global()->Get(v8::String::NewFromUtf8(isolate, "process"));
    v8::Local<v8::Function> process_function = v8::Local<v8::Function>::Cast(process_function_value);
    this->process_function = v8::Eternal<v8::Function>(this->isolate, process_function);

    printf("v8 started\n");

    // v8::Local<v8::Value>    query_function_value = local_context->Global()->Get(v8::String::NewFromUtf8(isolate, "query"));
    // v8::Local<v8::Function> query_function = v8::Local<v8::Function>::Cast(query_function_value);
    // this->query_function = v8::Eternal<v8::Function>(this->isolate, query_function);

    return true;
}

result<bool> ProcessV8Engine::process(std::vector<float *> samplesIn, std::vector<float *> samplesOut, std::vector<MIDIEvent> midiIn, std::vector<MIDIEvent> midiOut) {
    // Run the script to get the result.
    v8::HandleScope handle_scope(this->isolate);
    v8::Local<v8::Context> local_context = this->context.Get(this->isolate);
    v8::Context::Scope context_scope(local_context);

    v8::Local<v8::Array> jsSamplesIn = v8::Array::New(this->isolate, samplesIn.size());
    for (int i = 0; i < samplesIn.size(); i++) {
        v8::Local<v8::ArrayBuffer>  arrayBuffer  = v8::ArrayBuffer::New(this->isolate, samplesIn.at(i), this->samples_per_frame * sizeof(float));
        v8::Local<v8::Float32Array> float32Array = v8::Float32Array::New(arrayBuffer, 0, this->samples_per_frame);
        jsSamplesIn->Set(i, float32Array);
    }

    v8::Local<v8::Array> jsSamplesOut = v8::Array::New(this->isolate, samplesOut.size());
    for (int i = 0; i < samplesOut.size(); i++) {
        v8::Local<v8::ArrayBuffer>  arrayBuffer  = v8::ArrayBuffer::New(this->isolate, samplesOut.at(i), this->samples_per_frame * sizeof(float));
        v8::Local<v8::Float32Array> float32Array = v8::Float32Array::New(arrayBuffer, 0, this->samples_per_frame);
        jsSamplesOut->Set(i, float32Array);
    }

    v8::Local<v8::Array> jsMidiIn = v8::Array::New(this->isolate, 0);
    for (int i = 0; i < midiIn.size(); i++) {
        struct MIDIEvent& midiEvent = midiIn.at(i);
        v8::Local<v8::Object> jsMidiEvent = v8::Object::New(this->isolate);
        v8::Local<v8::ArrayBuffer> arrayBuffer = v8::ArrayBuffer::New(this->isolate, midiEvent.data, midiEvent.length * sizeof(char));
        v8::Local<v8::Uint8Array>  uint8Array  = v8::Uint8Array::New(arrayBuffer, 0, midiEvent.length);
        // jsMidiEvent->CreateDataProperty(local_context, v8::String::NewFromUtf8(this->isolate, "time"), v8::Number::New(this->isolate, (double)midiEvent.time));
        // jsMidiEvent->CreateDataProperty(local_context, v8::String::NewFromUtf8(this->isolate, "data"), uint8Array);

        jsMidiIn->Set(i, uint8Array);
    }

    v8::Local<v8::Value> args[] = {jsSamplesIn, jsSamplesOut, jsMidiIn, jsMidiIn};
    this->process_function.Get(this->isolate)->Call(local_context->Global(), 4, args);

    // Handle pending API query
    if (this->query_flag) {
        std::unique_lock<std::mutex> lk(this->query_mutex);

        this->query_response = "";

        v8::Local<v8::Value> query_args[] = {
            v8::String::NewFromUtf8(this->isolate, this->query_endpoint.c_str()),
            v8::String::NewFromUtf8(this->isolate, this->query_request.c_str())
        };

        v8::Handle<v8::Value> js_result = this->query_function.Get(this->isolate)->Call(local_context->Global(), 2, query_args);
        if (js_result->IsString()) {
            //todo - This should be simply - v8::String::Utf8Value result(js_result) - crashes with SEGV;
            v8::Handle<v8::String> v8String = v8::Handle<v8::String>::Cast(js_result);

            char buffer[2048];
            v8String->WriteUtf8(buffer, v8String->Utf8Length());
            std::string response(buffer, buffer + v8String->Utf8Length());
            this->query_response = response;
        }
        this->query_flag = false;

        lk.unlock();
        this->query_cv.notify_one();
    }

    return true;
}

result<bool> ProcessV8Engine::stop(void) {
    // Dispose the isolate and tear down V8.
    this->isolate->Dispose();
    v8::V8::Dispose();
    v8::V8::ShutdownPlatform();
    delete this->create_params.array_buffer_allocator;

    return true;
}

std::string ProcessV8Engine::query(std::string endpoint, std::string request) {
    this->query_endpoint = endpoint;
    this->query_request  = request;
    this->query_flag     = true;

    std::unique_lock<std::mutex> lk(this->query_mutex);
    this->query_cv.wait(lk);

    return this->query_response;
}
