// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type LiveStream struct {
	// Unique identifier for the Live Stream. Max 255 characters.
	Id string `json:"id,omitempty"`
	// Time the Live Stream was created, defined as a Unix timestamp (seconds since epoch).
	CreatedAt string `json:"created_at,omitempty"`
	// Unique key used for streaming to a Mux RTMP endpoint. This should be considered as sensitive as credentials, anyone with this stream key can begin streaming.
	StreamKey string `json:"stream_key,omitempty"`
	// The Asset that is currently being created if there is an active broadcast.
	ActiveAssetId string `json:"active_asset_id,omitempty"`
	// An array of strings with the most recent Assets that were created from this live stream.
	RecentAssetIds []string `json:"recent_asset_ids,omitempty"`
	// `idle` indicates that there is no active broadcast. `active` indicates that there is an active broadcast and `disabled` status indicates that no future RTMP streams can be published.
	Status string `json:"status,omitempty"`
	// An array of Playback ID objects. Use these to create HLS playback URLs. See [Play your videos](https://docs.mux.com/guides/video/play-your-videos) for more details.
	PlaybackIds      []PlaybackId       `json:"playback_ids,omitempty"`
	NewAssetSettings CreateAssetRequest `json:"new_asset_settings,omitempty"`
	// Arbitrary metadata set for the asset. Max 255 characters.
	Passthrough string `json:"passthrough,omitempty"`
	// The live stream only processes the audio track if the value is set to true. Mux drops the video track if broadcasted.
	AudioOnly bool `json:"audio_only,omitempty"`
	// When live streaming software disconnects from Mux, either intentionally or due to a drop in the network, the Reconnect Window is the time in seconds that Mux should wait for the streaming software to reconnect before considering the live stream finished and completing the recorded asset. **Min**: 0.1s. **Max**: 300s (5 minutes).
	ReconnectWindow float32 `json:"reconnect_window,omitempty"`
	// Latency is the time from when the streamer does something in real life to when you see it happen in the player. Set this if you want lower latency for your live stream. **Note**: Reconnect windows are incompatible with Reduced Latency and will always be set to zero (0) seconds. See the [Reduce live stream latency guide](https://docs.mux.com/guides/video/reduce-live-stream-latency) to understand the tradeoffs.
	ReducedLatency bool `json:"reduced_latency,omitempty"`
	// Latency is the time from when the streamer does something in real life to when you see it happen in the player. Setting this option will enable compatibility with the LL-HLS specification for low-latency streaming. This typically has lower latency than Reduced Latency streams, and cannot be combined with Reduced Latency. Note: Reconnect windows are incompatible with Low Latency and will always be set to zero (0) seconds.
	LowLatency bool `json:"low_latency,omitempty"`
	// Each Simulcast Target contains configuration details to broadcast (or \"restream\") a live stream to a third-party streaming service. [See the Stream live to 3rd party platforms guide](https://docs.mux.com/guides/video/stream-live-to-3rd-party-platforms).
	SimulcastTargets []SimulcastTarget `json:"simulcast_targets,omitempty"`
	// True means this live stream is a test live stream. Test live streams can be used to help evaluate the Mux Video APIs for free. There is no limit on the number of test live streams, but they are watermarked with the Mux logo, and limited to 5 minutes. The test live stream is disabled after the stream is active for 5 mins and the recorded asset also deleted after 24 hours.
	Test bool `json:"test,omitempty"`
}
