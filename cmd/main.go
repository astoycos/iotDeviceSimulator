package main


import (
	
	"time"
)

func main() { 

	for{
		<-time.After(2 * time.Second)
	getFrame("https://manifest.googlevideo.com/api/manifest/hls_playlist/expire/1582089400/ei/WHBMXrOzLafM8gSS9Y9Q/ip/144.121.20.163/id/1EiC9bvVGnk.1/itag/96/source/yt_live_broadcast/requiressl/yes/ratebypass/yes/live/1/goi/160/sgoap/gir%3Dyes%3Bitag%3D140/sgovp/gir%3Dyes%3Bitag%3D137/hls_chunk_host/r5---sn-ab5szn76.googlevideo.com/vprv/1/playlist_type/DVR/initcwndbps/25900/mm/44/mn/sn-ab5szn76/ms/lva/mv/m/mvi/4/pl/22/dover/11/keepalive/yes/fexp/23842630/mt/1582067719/disable_polymer/true/sparams/expire,ei,ip,id,itag,source,requiressl,ratebypass,live,goi,sgoap,sgovp,vprv,playlist_type/sig/ADKhkGMwRQIgMDhT8YoWRLOftmMFrypzb-CMn7sKaY-LV7_iWxxBUjACIQD1ahSS--dWbiXvqywScKeg8OyCZDZf4MWnQTnbUfS6WA%3D%3D/lsparams/hls_chunk_host,initcwndbps,mm,mn,ms,mv,mvi,pl/lsig/ABSNjpQwRgIhAI0VD8lNVUVXf4sSEghjHwDQNASArtzbYraWZk9z4J4xAiEAkUHcdi53mu2O1e09-EfTEHos7kDjwwwhyVGVmDgC_0Q%3D/playlist/index.m3u8")

	pushFrame("iot-http-adapter-enmasse-infra.apps.astoycos-ocp.shiftstack.com", "sensor1@myapp.iot", "hono-secret")

	<-time.After(20 * time.Second)
	}
}	