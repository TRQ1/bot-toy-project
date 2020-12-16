'use strict';

import puppeteer from 'puppeteer';
import dotenv from 'dotenv';
import util from 'util';
import dateFormat from 'dateformat'

dotenv.config();

const zabbix_user = process.env.ZABBIX_USER;
const zabbix_password = process.env.ZABBIX_PASSWORD;
const newrelic_user = process.env.NEWRELIC_USER;
const newrelic_password = process.env.NEWRELIC_PASSWORD;

const basic_zabbix = process.env.ZABBIX_URL
const basic_newrelic = process.env.NEWRELIC_URL

const urls = {
  zabbix: {
    login: ZABBIX_URL,
    global: util.format('%s%s', basic_zabbix, '1'),
  },
  newrelic: {
    login: NEWRELIC_URL,
    global: util.format('%s%s', basic_newrelic, '420907929'),
  }
}

const screenshot = async(target, board) => {
  const browser = await puppeteer.launch({
    executablePath: process.env.CHROME_BIN || null,
    args: ['--no-sandbox', '--headless', '--disable-gpu']
  });
  const page = await browser.newPage();
  const file_name = util.format('%s_%s_%s.png', target, board, dateFormat(Date.now(), 'yyyymmdd_HHMMss'))
   
  if(target == 'zabbix'){
      // zabbix login
      await page.goto(urls.zabbix.login, {waitUntil: 'networkidle2'});
      await page.type('#name', zabbix_user);
      await page.type('#password', zabbix_password);
      await page.click('#enter');

      // zabbix screenshot
      await page.goto(urls.zabbix[board], {waitUntil: 'networkidle2'});
      await page.screenshot({path: file_name, fullPage: true});

    } else if(target == 'newrelic'){
      await page.emulateTimezone('Asia/Seoul')
      await page.setViewport({width: 1350, height:1000, deviceScaleFactor: 1 });

      // newrelic login
      await page.goto(urls.newrelic.login, {waitUntil: 'load'});
      await page.type('#login_email', newrelic_user);
      await page.type('#login_password', newrelic_password);
      await page.click('#login_submit');

      // newrelic screenshot
      await page.goto(urls.newrelic[board], {waitUntil: 'load'});
      await page.waitFor(3000);
      await page.screenshot({path: file_name, clip:{x: 0, y: 0, width: 1350, height: 1000 }});
    } else {
      // for test
      await page.goto('https://www.naver.com', {waitUntil: 'networkidle2'});
      await page.screenshot({path: 'test.png', fullPage: true});
    }

    await browser.close()

    return file_name
  };

   export { screenshot, urls };
