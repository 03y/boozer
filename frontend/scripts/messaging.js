function setMessage(message, isError = false) {
    const messageContainer = document.getElementById('message-container');
    const messageElement = document.createElement('div');
    messageElement.textContent = message;
    messageContainer.innerHTML = '';

    if (isError) {
        messageElement.className = 'error';
    } else {
        messageElement.className = 'success';
    }

    messageContainer.appendChild(messageElement);
}
