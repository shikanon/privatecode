from fastapi import FastAPI,Request

app = FastAPI()


@app.get("/")
def read_root(request: Request):
    host = request.headers.get('host')
    print(host)
    authority = request.headers.get('authority')
    print(authority)
    return {"message": "Hello World", "host": host, "authority": authority}

@app.get("/{version}")
def read_root(request: Request,version: str):
    host = request.headers.get('host')
    print(host)
    authority = request.headers.get('authority')
    print(authority)
    print(version)
    return {"message": "get the version api", "version": version, "host": host, "authority": authority}