#include "Process.hpp"

#include "include/libplatform/libplatform.h"
#include "include/v8.h"

#include <atomic>
#include <mutex>
#include <condition_variable>

class ProcessV8Engine : public Process {
  public:
    ProcessV8Engine() : query_flag(0) {}
    virtual result<bool> init(void);
    virtual result<bool> start(int sampling_rate, int samples_per_frame);
    virtual result<bool> process(std::vector<float *> samplesIn, std::vector<float *> samplesOut, std::vector<MIDIEvent> midiIn, std::vector<MIDIEvent> midiOut);
    virtual result<bool> stop(void);
    virtual std::string query(std::string endpoint, std::string request);
  private:
    void compile(std::string filename);
    void compile_source(std::string filename);

    int samples_per_frame;
    v8::Isolate::CreateParams create_params;
    v8::Isolate* isolate;
    v8::Eternal<v8::Context>  context;
    v8::Eternal<v8::Function> process_function;
    v8::Eternal<v8::Function> query_function;

    std::atomic<bool>       query_flag;
    std::mutex              query_mutex;
    std::condition_variable query_cv;
    std::string             query_endpoint;
    std::string             query_request;
    std::string             query_response;
};
