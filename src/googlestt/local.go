package googlestt

import (
        "context"
        "fmt"
        "io/ioutil"
        "log"
	"io"
	"os"
        speech "cloud.google.com/go/speech/apiv1"
        speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func send(w io.Writer, client *speech.Client, filename string) error {
        ctx := context.Background()
        data, err := ioutil.ReadFile(filename)
        if err != nil {
                return err
        }
	log.Println("Length of audio bytes array:", len(data))
        // Send the contents of the audio file with the encoding and
        // and sample rate information to be transcripted.
        req := &speechpb.LongRunningRecognizeRequest{
                Config: &speechpb.RecognitionConfig{
                        Encoding:        speechpb.RecognitionConfig_LINEAR16,
                        //SampleRateHertz: 16000,
			EnableWordTimeOffsets: true,
                        LanguageCode:    "en-US",
                },
                Audio: &speechpb.RecognitionAudio{
                        AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
                },
        }

        op, err := client.LongRunningRecognize(ctx, req)
        if err != nil {
                return err
        }
        resp, err := op.Wait(ctx)
        if err != nil {
                return err
        }

        // Print the results.
        for _, result := range resp.Results {
		fmt.Println("Here")
                for _, alt := range result.Alternatives {
                        fmt.Fprintf(w, "\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
                }
        }
        return nil
}

func MakeLocalRequest(audio_path string) {
	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create clinet: %v", err)
	}
	err = send(os.Stdout, client, audio_path)
	if err != nil {
		log.Fatalf("Failed to create clinet: %v", err)
	}
}
