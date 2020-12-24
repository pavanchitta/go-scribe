package googlestt

import (
        "context"
        "fmt"
        "log"
	"io"
	"os"
        speech "cloud.google.com/go/speech/apiv1"
        speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func sendGCS(w io.Writer, client *speech.Client, gcsURI  string) error {
        ctx := context.Background()
        // Send the contents of the audio file with the encoding and
        // and sample rate information to be transcripted.
        req := &speechpb.LongRunningRecognizeRequest{
                Config: &speechpb.RecognitionConfig{
                        Encoding:        speechpb.RecognitionConfig_LINEAR16,
                        //SampleRateHertz: 16000,
                        LanguageCode:    "en-US",
                },
                Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Uri{Uri: gcsURI},
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
                for _, alt := range result.Alternatives {
                        fmt.Fprintf(w, "\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
                }
        }
        return nil
}

func MakeRemoteRequest(gcsURI string) {
	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create clientt: %v", err)
	}
	err = sendGCS(os.Stdout, client, gcsURI)
	if err != nil {
		log.Fatalf("Failed to perform Long running Request: %v", err)
	}
}
