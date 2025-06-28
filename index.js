import { Octokit } from "@octokit/core";
import {createServer} from 'http';

const RAW_URL = "https://raw.githubusercontent.com/keshuook/keshuook.github.io";
const contentTypeTable = {
    'html': 'text/html',
    'js': 'text/javascript',
    'css': 'text/css',
    'gif': 'image/gif',
    'jpg': 'image/jpeg',
    'png': 'image/png',
    'mp3': 'audio/mpeg',
    'ogg': 'audio/ogg',
    'json': 'application/json',
    'ttf': 'font/ttf',
    'svg': 'image/svg+xml',
    'txt': 'text/plain',
    'ico': 'image/vnd.microsoft.icon'
}

const octokit = new Octokit({
    auth: process.env.TOKEN
});

let reqData = [];
let page = 1;
while(page < 20) {
    console.log(`Request for page ${page} made.`);
    const res = await octokit.request("GET /repos/keshuook/keshuook.github.io/commits", {
        per_page: 100,
        page: page++
    });
    reqData =  reqData.concat(res.data);
    if(res.data.length == 0) break;
}
const shaData = reqData.map(v => {
    return {sha: v.sha, date: (new Date(v.commit.committer.date)).valueOf()};
});

createServer(async (req, res) => {
    const paths = req.url.split("/");    
    const target = (new Date(parseInt(paths[1]), parseInt(paths[2]) - 1, parseInt(paths[3]))).valueOf();
    if(isNaN(target)) return;

    if(req.url.endsWith("/")) {
        paths.pop();
        paths.push("index.html");
    } else if(paths[paths.length - 1].search(".") == -1) {
        paths[paths.length - 1] = paths[paths.length - 1].concat(".html");
    }

    const type = paths[paths.length - 1].split(".").pop();

    let L = 0,U = shaData.length,M;

    while((U - L) > 1) {
        M = Math.floor(((L+U)/2));
        console.log(`L: ${L}, M: ${M}, U: ${U}, D: ${shaData[M].date}`);
        console.log(U - L);
        if(shaData[M].date > target) {
            L = M;
        } else {
            U = M;
        }
    }

    console.log(paths);
    const data = await fetch(`${RAW_URL}/${shaData[M].sha}/${paths.slice(4).join("/")}`);
    console.log(`${RAW_URL}/${shaData[M].sha}/${paths.slice(4).join("/")}`);
    let mimeType = contentTypeTable[type];
    if(!mimeType) mimeType = "application/octet-stream";

    res.writeHead(data.status, {"content-type": mimeType});
    res.end(Buffer.from(await data.arrayBuffer()));
}).listen(process.env.PORT || 80);
