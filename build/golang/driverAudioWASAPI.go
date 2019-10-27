// +build windows

package nf

/*
#cgo windows,386 LDFLAGS: -lole32 -loleaut32
#cgo windows,amd64 LDFLAGS: -lole32 -loleaut32

#define INITGUID
#include <windows.h>
#include <mmdeviceapi.h>
#include <audioclient.h>
#include <stdio.h>

#define REFTIMES_PER_SEC  1000000
#define REFTIMES_PER_MILLISEC  10000

#define EXIT_ON_ERROR(hres)  \
              printf("result %d\n", hr); if (FAILED(hres)) { goto Exit; }
#define SAFE_RELEASE(punk)  \
              if ((punk) != NULL)  \
                { (punk)->lpVtbl->Release(punk); (punk) = NULL; }

#define MAX_CHANNELS 8
#define MAX_SAMPLES  48000

typedef struct {
	HANDLE thread;
	HANDLE renderEvent;
	IAudioClient *pAudioClient;
	IAudioRenderClient *pRenderClient;
	UINT32 bufferFrameCount;

	int channel_in_count;
	int channel_out_count;
	float* channel_in_float32[MAX_CHANNELS];
	float* channel_out_float32[MAX_CHANNELS];
} wasapi_c_client;

wasapi_c_client client;

extern void goAudioWASAPICallback(void *arg, int blockLength,
	int channelInCount, void *channelIn,
	int channelOutCount, void *channelOut);

static inline DWORD WINAPI runWASAPIThread(void* arg) {
	UINT32 numFramesPadding;
	UINT32 numFramesAvailable;
	BYTE *pData;
	HRESULT hr;
    UINT32 bufferFrameCount;
    DWORD flags = 0;

	printf("start thread\n");

	while (1) {
		// See how much buffer space is available.
	    client.pAudioClient->lpVtbl->GetCurrentPadding(client.pAudioClient, &numFramesPadding);

		numFramesAvailable = client.bufferFrameCount - numFramesPadding;

		client.pRenderClient->lpVtbl->GetBuffer(client.pRenderClient, numFramesAvailable, &pData);

		for (int i = 0; i < client.channel_in_count; i++) {
			memset(client.channel_in_float32[i], 0, numFramesAvailable * sizeof(float));
		}

		goAudioWASAPICallback(arg, numFramesAvailable,
			client.channel_in_count, client.channel_in_float32,
			client.channel_out_count, client.channel_out_float32
		);

		int channel = 0;
		int sample  = 0;
		int samples = numFramesAvailable * client.channel_out_count;
		for (int i = 0; i < samples; i++) {
			float *ps = (float *)(pData + (i * sizeof(float)));
			*ps = client.channel_out_float32[channel][sample];
			channel++;
			if (channel >= client.channel_out_count) {
				sample++;
				channel = 0;
			}
		}
		client.pRenderClient->lpVtbl->ReleaseBuffer(client.pRenderClient, numFramesAvailable, 0);
		WaitForSingleObject( client.renderEvent, INFINITE );
	}
}

static inline wasapi_c_client* gowasapi_client_open(uintptr_t arg) {
	HRESULT hr;
	IMMDeviceEnumerator *pEnumerator = NULL;
	IMMDevice *pDevice = NULL;
	WAVEFORMATEX *pwfx = NULL;
	REFERENCE_TIME hnsRequestedDuration = REFTIMES_PER_SEC;
	REFERENCE_TIME hnsActualDuration;

	CoInitialize(NULL);

	printf("CoCreateInstance\n");
	hr = CoCreateInstance(
	   	&CLSID_MMDeviceEnumerator, NULL,
	   	CLSCTX_ALL, &IID_IMMDeviceEnumerator,
	   	(void**)&pEnumerator);
		// printf("result %d\n", hr);
	EXIT_ON_ERROR(hr)

	printf("GetDefaultAudioEndpoint\n");
	hr = pEnumerator->lpVtbl->GetDefaultAudioEndpoint(pEnumerator,
					eRender, eConsole, &pDevice);
	EXIT_ON_ERROR(hr)

	printf("Activate\n");
	hr = pDevice->lpVtbl->Activate(pDevice,
				&IID_IAudioClient, CLSCTX_ALL,
				NULL, (void**)&client.pAudioClient);
	EXIT_ON_ERROR(hr)

	printf("GetMixFormat\n");
	hr = client.pAudioClient->lpVtbl->GetMixFormat(client.pAudioClient, &pwfx);
	client.channel_in_count  = pwfx->nChannels;
	client.channel_out_count = pwfx->nChannels;
	EXIT_ON_ERROR(hr)

	hr = client.pAudioClient->lpVtbl->Initialize(client.pAudioClient,
					 AUDCLNT_SHAREMODE_SHARED,
					 AUDCLNT_STREAMFLAGS_EVENTCALLBACK,
					 hnsRequestedDuration,
					 0,
					 pwfx,
					 NULL);
	EXIT_ON_ERROR(hr)

	hr = client.pAudioClient->lpVtbl->GetBufferSize(client.pAudioClient, &client.bufferFrameCount);
	EXIT_ON_ERROR(hr)

	hr = client.pAudioClient->lpVtbl->GetService(client.pAudioClient,
					 &IID_IAudioRenderClient,
					 (void**)&client.pRenderClient);
	EXIT_ON_ERROR(hr)

	client.renderEvent = CreateEvent(NULL, FALSE, FALSE, NULL);
	hr = client.pAudioClient->lpVtbl->SetEventHandle(client.pAudioClient, client.renderEvent);
	EXIT_ON_ERROR(hr)

	hr = client.pAudioClient->lpVtbl->Start(client.pAudioClient);
	EXIT_ON_ERROR(hr)

	for (int i = 0; i < client.channel_in_count; i++) {
		client.channel_in_float32[i] = calloc(MAX_SAMPLES, sizeof(float));
	}

	for (int i = 0; i < client.channel_out_count; i++) {
		client.channel_out_float32[i] = calloc(MAX_SAMPLES, sizeof(float));
	}

	client.thread = CreateThread(NULL, 0, runWASAPIThread, (void *)arg, CREATE_SUSPENDED, NULL);
	// SetThreadPriority((void*) client.thread, THREAD_PRIORITY_HIGHEST);
	ResumeThread((void*) client.thread);

	Exit:
		CoTaskMemFree(pwfx);
		SAFE_RELEASE(pEnumerator)
		SAFE_RELEASE(pDevice)
}

static inline void gowasapi_client_close() {
	SAFE_RELEASE(client.pAudioClient)
	SAFE_RELEASE(client.pRenderClient)
}


*/
import "C"

import (
	"unsafe"

	"github.com/jacoblister/noisefloor/app/audiomodule"
)

type driverAudioWASAPI struct {
	audioProcessor audiomodule.AudioProcessor
	driverMidi     driverMidi
}

//export goAudioWASAPICallback
func goAudioWASAPICallback(arg unsafe.Pointer, blockLength C.int,
	channelInCount C.int, channelIn unsafe.Pointer,
	channelOutCount C.int, channelOut unsafe.Pointer) {

	driverAudio := (*driverAudioWASAPI)(arg)

	goAudioCallback(driverAudio, int(blockLength), int(channelInCount), channelIn, int(channelOutCount), channelOut)
}

func (d *driverAudioWASAPI) getDriverMidi() driverMidi {
	return d.driverMidi
}

func (d *driverAudioWASAPI) setDriverMidi(driverMidi driverMidi) {
	d.driverMidi = driverMidi
}

func (d *driverAudioWASAPI) getAudioProcessor() audiomodule.AudioProcessor {
	return d.audioProcessor
}

func (d *driverAudioWASAPI) setAudioProcessor(audioProcessor audiomodule.AudioProcessor) {
	d.audioProcessor = audioProcessor
}

func (d *driverAudioWASAPI) start() {
	uintPtr := uintptr(unsafe.Pointer(d))
	C.gowasapi_client_open((C.ulonglong)(uintPtr))
}

func (d *driverAudioWASAPI) stop() {
	C.gowasapi_client_close()
}

func (d *driverAudioWASAPI) samplingRate() int {
	return 48000
}
