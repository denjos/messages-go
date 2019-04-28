let url=`ws://${window.location.host}/ws`;
let ws=new WebSocket(url);
let container=document.getElementById("container");
ws.onmessage=function (msg) {
    let newMessage=JSON.parse(msg.data);
    let div=document.createElement('div');
    div.innerText=newMessage.content;
    //container.innerHTML=`<pre>${msg.data}</pre>`;
    container.appendChild(div);
};

function requestAjax(method, url, obj) {
    return new Promise(function (resolve, reject) {
        let xhr = new XMLHttpRequest();
        xhr.open(method, url, true);
        xhr.setRequestHeader('Content-Type', 'application/json');
        if (sessionStorage.getItem('token')) {
            xhr.setRequestHeader('Authorization', sessionStorage.getItem('token'))
        }
        xhr.addEventListener('load', (e) => {
            let self = e.target;
            let result = {
                status: self.status,
                response: JSON.parse(self.response),
            };
            resolve(result);
        });
        xhr.addEventListener('error', (e) => {
            let self = e.target;
            console.log(self);
            reject(self);
        });
        xhr.send(obj);
// Promises
    });
};

function $(el) {
    return document.querySelector(el);
}