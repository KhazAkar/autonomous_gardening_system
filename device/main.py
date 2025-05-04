from falcon.asgi import App
import falcon

class HelloWorld:
    async def on_get(self, req, resp):
        resp.media = {'message': 'Hello, world!'}
        resp.content_type = 'application/json'
        resp.status = falcon.HTTP_200

app = App()
app.add_route('/', HelloWorld())
