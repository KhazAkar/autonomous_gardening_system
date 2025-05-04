
import falcon

class DeviceResource:
    def on_get(self, req, resp):
        resp.status = falcon.HTTP_200
        resp.body = 'Hello from device!'


if __name__ == "__main__":
    main()
