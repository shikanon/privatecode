// 导入包
const puppeteer = require('puppeteer');

const Koa = require('koa');
const app = new Koa();

// 存储browserWSEndpoint列表
let WSE_LIST = [];

// browser 初始化，将bwse存储复用
(async() =>{
  // 因为服务器内核不支持sandbox，所以只能启用--no-sandbox
  const browser = await puppeteer.launch({args: ['--no-sandbox', '--disable-setuid-sandbox','--no-first-run']});
  // 存储节点以便能重新连接到 Chromium
  const browserWSEndpoint = await browser.wsEndpoint();
  WSE_LIST = [browserWSEndpoint]
})();

app.use(async ctx =>{
  console.log(ctx.url);
  let time1 = new Date().getTime();
  let url = 'https://developer.orbbec.com.cn' + ctx.url
  console.log(url);

  // 恢复节点
  let browserWSEndpoint = WSE_LIST[0]
  console.log(browserWSEndpoint)
  const browser = await puppeteer.connect({browserWSEndpoint});
  
  // 开启新的标签页
  let page = await browser.newPage();
  await page.setJavaScriptEnabled(true);
  // 由于只关心渲染后的dom树，所以对css，font，image等都做了屏蔽
  await page.setRequestInterception(true); 
  page.on('request', (req) => {
    if(req.resourceType() == 'stylesheet' || req.resourceType() == 'font' || req.resourceType() == 'image'){
            req.abort();
        }
        else {
            req.continue();
        }
    });


  // waitUntil 主要包括四个值，'load','domcontentloaded','networkidle2','networkidle0'
  // 分别表示在xx之后才确定为跳转完成
  // load - 页面的load事件触发时
  // domcontentloaded - 页面的 DOMContentLoaded 事件触发时
  // networkidle2 - 只有2个网络连接时触发（至少500毫秒后）
  // networkidle0 - 不再有网络连接时触发（至少500毫秒后）
  await page.goto(url, { waitUntil: ['load','domcontentloaded','networkidle2'] });

  ctx.body = await page.content();
  // 关闭标签页
  await page.close();

  // 断开连接
  await browser.disconnect();
  
  let time2 = new Date().getTime();
  console.log((time2-time1)/1000)
  console.log("finish");
  // 关闭浏览器
  // await browser.close();
});

app.listen(8000);
