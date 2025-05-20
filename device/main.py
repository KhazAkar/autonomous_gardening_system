from falcon.asgi import App
from falcon import Response
import falcon


class HelloWorld:
    async def on_get(self, req, resp: Response):
        resp.media = {"message": "Hello, world!"}
        resp.content_type = "application/json"
        resp.status = falcon.HTTP_200

    async def on_post(self, req, resp: Response):
        resp.media = {"message": "Hello, world!"}
        resp.content_type = "application/json"
        resp.status = falcon.HTTP_200


app = App()
app.add_route("/", HelloWorld())

