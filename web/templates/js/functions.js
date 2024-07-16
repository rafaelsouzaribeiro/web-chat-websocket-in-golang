
const chat = document.getElementById('spacer');
const messageInput = document.getElementById('message');
const usernameInput = document.getElementById('username');
const messages = document.getElementById('spacer-ms')
const messagesScroll = document.getElementById('messages')
const scrolldown = document.getElementById('scroll-down')
const load = document.getElementById('load')
let socket;
let hasMoreMessages = true;
let hasMoreUsers=true
let startIndex = -20;
let startmessageIndex = -20;


window.addEventListener('beforeunload', function (event) {
    if(socket){
       socket.close();
    }

    event.preventDefault();                
        
});

scrolldown.addEventListener('click',function(event){
    messagesScroll.scrollTo(0, messagesScroll.scrollHeight);
});

chat.addEventListener('scroll', function(event) {
    if (chat.scrollTop === 0 && hasMoreUsers) {
        loadPreviousUsers();
    }
});

load.addEventListener('click',function(event){
    if (hasMoreUsers) {
        loadPreviousUsers();
    }
});


function updateVisibleMessages() {
    let children = messages.getElementsByClassName("horadate");
    let visibleValues = "";

    for (let index = 0; index < children.length; index++) {
        const element = children[index];
        const rect = element.getBoundingClientRect();

        if (rect.top >= 0 && (rect.bottom+50) <= window.innerHeight) {
            document.getElementById('tempo').innerHTML=element.innerHTML; 
        }

    }
     
}

messagesScroll.addEventListener('scroll', function(event) {
    if (messagesScroll.scrollTop === 0 && hasMoreMessages) {
        loadPreviousMessages();
    }

    updateVisibleMessages()
});

function GetTime(v){
    let currentTime = new Date(v);
    let hours = currentTime.getHours();
    let minutes = currentTime.getMinutes();

    return formattedTime = `${String(hours).padStart(2, '0')}:${String(minutes).padStart(2, '0')}`;
}


function loadPreviousUsers() {
    fetch(`/last-users/${startIndex}`)
        .then(response => response.json())
        .then(data => {
            if(data.messages!=null){
                data.messages.reverse().forEach(msg => {
                    console.log(msg)
                    const messageElement = document.createElement('div');
                    messageElement.classList.add('message');
                    messageElement.innerHTML = `${msg.username}: ${msg.message}`;
                    chat.insertBefore(messageElement, chat.firstChild);
                });
                startIndex -= 20;

                const childDivs = chat.getElementsByTagName('div');
                
                if(childDivs.length>19){
                    chat.scrollTop = chat.scrollHeight;
                }
                hasMoreUsers = data.hasMore;
        }

        })
        .catch(error => console.error('Error fetching previous messages:', error));
}

function loadPreviousMessages() {
    fetch(`/last-messages/${startmessageIndex}`)
        .then(response => response.json())
        .then(data => {
            if(data.messages!=null){
                data.messages.reverse().forEach(msg => {
                    console.log(msg)
                    formattedTime=GetTime(msg.time)

                    const messageElement = document.createElement('div');
                    messageElement.classList.add('message');
                    messageElement.innerHTML = `${msg.username}: ${msg.message}`;
                    v=formatTime(msg.time);
                    messages.insertBefore(messageElement, messages.firstChild);

                    const horadate = document.createElement('div');
                    horadate.classList.add('horadate');
                    horadate.innerHTML = `${v} - ${formattedTime}`;
                    messageElement.appendChild(horadate)
                    
                                              
                });
                startmessageIndex -= 20;

                const childDivs = messages.getElementsByTagName('div');
                
                if(childDivs.length>19){
                    messagesScroll.scrollTop = messages.scrollHeight;
                }
                hasMoreMessages = data.hasMore;
                updateVisibleMessages()
         }

        })
        .catch(error => console.error('Error fetching previous messages:', error));
}

function connect() {
    const username = usernameInput.value;
    if (!username) {
        alert('Please enter a username.');
        return;
    }

    
    if (socket) {
        chat.innerHTML=""
        socket.close();
    }

    socket = new WebSocket(webSocketURL);

    socket.onopen = function() {
        console.log('Connected to the server');
        sendMessage('connected', 'connect');
    };

    socket.onmessage = function(event) {
        const msg = JSON.parse(event.data);

        if (msg.username && msg.message) {

            const messageElement = document.createElement('div');
            messageElement.classList.add('message');
            messageElement.innerHTML = `${msg.username}: ${msg.message}`;
            let v = formatTime(msg.time);
               
            if (msg.type=="message"){
                formattedTime=GetTime(msg.time)
                const horadate = document.createElement('div');
                horadate.classList.add('horadate');
                horadate.innerHTML = `${v} - ${formattedTime}`;
                messageElement.appendChild(horadate);

                messages.appendChild(messageElement);

            }else{
               chat.appendChild(messageElement);
            }


            const childDivs = chat.getElementsByTagName('div');

            if(childDivs.length>13){
                chat.scrollTop = chat.scrollHeight;
            }

        }
    };

    socket.onclose = function() {
        chat.innerHTML=""
        messages.innerHTML=""
        console.log('Disconnected from the server');
    };

    socket.onerror = function(error) {
        console.log('WebSocket error: ' + error);
    };
}

function sendMessage(message, type) {
    if (!message) {
        message = messageInput.value;
    }
    const username = usernameInput.value;
    const currentTime = new Date().toISOString();
    if (socket && socket.readyState === WebSocket.OPEN) {
        socket.send(JSON.stringify({ Username: username, Message: message, Type: type,Time:currentTime }));
       // messageInput.value = '';
    } else {
        alert('WebSocket is not connected.');
    }
}


const emojiContainer = document.querySelector(".emoji-container");
const emojiInput = document.querySelector(".emoji-input");

const emojis = ["😀", "😃", "😄", "😁", "😆", "😅", "😂", "🤣", "😊", 
"😇","💗","💔","❤️‍🔥","❤","😍","😴","😌","😌","🤤","😱","😭","😩","🤬","🤡","👹","👺","👻","👽"
,"👾","🙌","🤝","🙏","👍","👎"];

emojis.forEach((emoji) => {
    const emojiDiv = document.createElement("div");
    emojiDiv.classList.add("emoji");
    emojiDiv.innerText = emoji;
    emojiContainer.appendChild(emojiDiv);
});

emojiInput.addEventListener("focus", () => {
    emojiContainer.style.display = "block";
});

emojiInput.addEventListener("blur", () => {
    setTimeout(() => {
    emojiContainer.style.display = "none";
    }, 200);
});

emojiContainer.addEventListener("click", (e) => {
    emojiInput.value += e.target.innerText;
});


function formatTime(time) {
    const today = new Date();
    const date = new Date(time);

    const isToday = today.toDateString() === date.toDateString();
    if (isToday) {
        return "Hoje";
    }

    const yesterday = new Date(today);
    yesterday.setDate(today.getDate() - 1);
    const isYesterday = yesterday.toDateString() === date.toDateString();
    if (isYesterday) {
        return "Ontem";
    }

    return date.toLocaleDateString();
}


usernameInput.addEventListener("keydown", (e) => {
    if ((e.keyCode == 32) && (e.ctrlKey || e.metaKey)) {
        
        console.log(e.keyCode)
        return;
    }
});