# coding=utf-8
'''
# Author: shikanon (shikanon@tensorbytes.com)
# File Created Time: 2020-03-25 2:21:35
# 
# Project: FastAPI-Project-Template
# File: main.py
# Description: 
# 
'''

import time
from fastapi import FastAPI, Request
from pydantic import BaseModel

from settings.config import Config
from model.user import User
from model.database import Database

app = FastAPI()

conf = Config()
conf.parse("settings/config.ini")
db = Database(conf.mysqldb)

class ModelUser(BaseModel):
    name: str
    passwd: str
    email: str
    usertype: str
    organization: str



@app.middleware("http")
async def auth(request: Request, call_next):
    start_time = time.time()
    response = await call_next(request)
    process_time = time.time() - start_time
    response.headers["X-Process-Time"] = str(process_time)
    return response


@app.get("/login")
async def login(username: str, passwd: str):
    record = db.session.query(User).filter_by(name=username).first()
    print(record)
    if not record:
        return {"result": "用户不存在"}
    return {"token": "xxxxxxx"}


@app.post("/user/add")
async def add_user(user: ModelUser):
    return {"result": "login success!"}