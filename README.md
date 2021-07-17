# **ATracker**
Aircraft tracking and notification client. Track all aircraft within the
set cordinate rectangular bounds and notifiy when one of the filtered aircraft is present.

## **Config**
From the template `config_template.yaml`, a copy of it needs to be made and named `config.yaml`.
Once that is done, the following config options could be set:
- **radarBounds.topRightCords.latitude=** latitude cordinate for top right edge of the bounds.
- **radarBounds.topRightCords.longitude=** longitude cordinate for top right edge of the bounds.
- **radarBounds.bottomLeftCords.latitude=** latitude cordinate for bottom left edge of the bounds.
- **radarBounds.bottomLeftCords.longitude=** longitude cordinate for bottom left edge of the bounds.
- **aircraft** list of all special aircrafts to be filtered. All aircrafts in this list must be designated by their ICAO code.
- **redis.addr=** the redis server address which will be used for caching flight ids.
- **redis.port=** the redis server port which will be used for caching flight ids.

## **Caching**
To avoid duplicate notifications for the same aircraft, each aircraft, by their fid,
is assisted a TTL of 5 minutes once detected. This caching is done via Redis with the TTL
currently not user configurable.

## **Push Notification**
As the software currently sits, the "push notification" is done via standard output. Plans to
extend this are in the works as I am working on implementing a universal push notification service
with atracker being one of the supported platforms.

## **Captcha Prevention**
As discovered via R/D of the internal API that I use to retrieve aircraft data, frequest aircraft
info requests result in Cloudflare sending captchas to solve. To prevent this, scans of the bounds
only occur every 10 seconds. Similarly, the number of concurrent aircraft workers goroutines is set to 5. Within each worker, a random timeout between [0,10] seconds is set. Thus, as a result, there can be 
delay in the notifiaction. Until further optimization is done to lower these values, these values
have been upperbounded to higher values to avoid high congestion resulting in temp bans and captchas. 

## **Installation & Runtime**
The following steps must be done for the tool to run:
1. Install Go v1.16.
2. Clone repo and install dependencies with `go mod download`.
3. Configure the filter and the bounds of the tracker as designated in [config](#config).
4. Two ways of running the tool:
   1. Either build the tool with `go build` and then run it was `./atracker` *Note that* `config.yaml` *must be present in same directory as* `atracker` *binary*.
   2. Run without build via `go run main.go`.
