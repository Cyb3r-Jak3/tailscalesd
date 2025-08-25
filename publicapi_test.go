package tailscalesd

import (
	"net/url"
	"testing"
)

func apiBaseForTest(tb testing.TB, surl string) string {
	tb.Helper()
	u, err := url.Parse(surl)
	if err != nil {
		tb.Fatal(err)
	}
	return u.Host
}

//func TestPublicAPIDiscovererDevices(t *testing.T) {
//	var wantPath = "/api/v2/tailnet/testTailnet/devices"
//	for tn, tc := range map[string]struct {
//		responder func(w http.ResponseWriter)
//		wantErr   error
//		want      []Device
//	}{
//		"returns failed request error when the server responds unsuccessfully": {
//			responder: func(w http.ResponseWriter) {
//				w.WriteHeader(http.StatusInternalServerError)
//			},
//			wantErr: errFailedAPIRequest,
//		},
//		"returns failed request error when the server responds with bad payload": {
//			responder: func(w http.ResponseWriter) {
//				w.Header().Set("Content-Type", "text/plain")
//				fmt.Fprintln(w, "This is decidedly not JSON.")
//			},
//			wantErr: errFailedAPIRequest,
//		},
//		"returns devices when the server responds with valid JSON": {
//			responder: func(w http.ResponseWriter) {
//				w.Header().Set("Content-Type", "application/json; encoding=utf-8")
//				_, _ = w.Write([]byte(`{"devices": [{"hostname":"testhostname","os":"beos"}]}`))
//			},
//			want: []Device{
//				{
//					Hostname: "testhostname",
//					OS:       "beos",
//					Tailnet:  "testTailnet",
//				},
//			},
//		},
//	} {
//		t.Run(tn, func(t *testing.T) {
//			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//				if got, want := r.URL.Path, wantPath; got != want {
//					t.Errorf("Devices: request URL path mismatch: got: %q want: %q", got, want)
//				}
//				tc.responder(w)
//			}))
//			defer server.Close()
//
//			d := PublicAPI("testTailnet", "testToken", WithHTTPClient(server.Client()), WithAPIHost(apiBaseForTest(t, server.URL)))
//			got, err := d.Devices(context.TODO())
//			if got, want := err, tc.wantErr; !errors.Is(got, want) {
//				t.Errorf("Devices: error mismatch: got: %q want: %q", got, want)
//			}
//			// Ignore the API field, which will be set to the arbitrary test
//			// server's host:port.
//			if diff := cmp.Diff(got, tc.want, cmpopts.IgnoreFields(Device{}, "API")); diff != "" {
//				t.Errorf("PublicAPI: mismatch (-got, +want):\n%v", diff)
//			}
//		})
//	}
//}
