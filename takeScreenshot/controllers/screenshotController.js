import { screenshot, urls } from '../screenShot.js';
import uploadFile from '../s3.js';
import axios from 'axios';
import util from 'util';

export const getMain = (req, res) => {
    return res.json({result: 'main success'})
};

export const postScreenshot = async(req, res) => {
    try {
        // slack slash command를 통해서 실행되지 않음
        if(req.body.token != process.env.SLACK_SLASH_COMMAND_TOKEN){
            return res.json({'result': '잘못된 경로를 통해 실행되었습니다.'})
        }

        const params = req.body.text.split(' ');
        const target = params[0]
        const board = params[1]
        const user_name = req.body.user_name
        
        var params_list = {};
        if(urls[target] && urls[target][board]){
            res.json({'result': util.format('%s에서 %s의 지표가 전송됩니다.', target, board)})
            let file_name = await screenshot(target, board)
            await uploadFile(file_name)
            await sleep(3000);
            await axios.get(util.format(AWS_URL+'imageName=%s&service=%s&board=%s&user=%s', file_name, target, board, user_name))
            .catch(function(err){
                console.log(err)
            })
        } else {
            // 잘못된 target 및 board 입력
            Object.keys(urls).forEach(e => {
                params_list[e] = Object.keys(urls[e]).filter(e => e != 'login')
            });
            return res.json({'result': '잘못 입력하였습니다. 다음 중에서 선택해주세요.\n' + JSON.stringify(params_list)})
        }
    } catch(e) {
        return res.json({result: 'fail', message: e.message})
    }
};

function sleep(ms) {
    return new Promise((resolve) => {
        setTimeout(resolve, ms);
    });
} 