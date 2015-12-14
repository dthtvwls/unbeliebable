package unbeliebable

import "net/http"

type Player struct {
	Playlist *Playlist
}

func (m *Player) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if song, err := m.Playlist.Shift(); err != nil {
		w.Write([]byte(`<!DOCTYPE html>
			<html>
				<head>
					<title>!</title>
				</head>
				<body>
					<p>No song available.</p>
					<script>setTimeout(location.reload.bind(location), 1000);</script>
				</body>
			</html>
		`))
	} else {
		w.Write([]byte(`<!DOCTYPE html>
			<html>
				<head>
					<title>▶</title>
				</head>
				<body>
					<script>
						window.onYouTubeIframeAPIReady = function () {
							var player = new YT.Player(document.body, {
								videoId: "` + song.ID + `",
								playerVars: { autoplay: 1 },
								events: {
									onReady: function (event) {
										event.target.playVideo();
										setInterval(function () {
											var pct = Math.round(player.getCurrentTime() / player.getDuration() * 100);
											console.log(pct);

											var xhr = new XMLHttpRequest();
											xhr.open("POST", "http://localhost/elapsedtime");
        									xhr.send(pct);
										}, 1000);
									},
									onStateChange: function (event) {
										if (event.data === YT.PlayerState.ENDED) location.reload();
									}
								}
							});
						};
						var script = document.createElement("script");
						script.src = "//www.youtube.com/iframe_api";
						document.body.appendChild(script);
					</script>
				</body>
			</html>
		`))
	}
}
