"use strict";
export const chatBody = document.querySelector(".chat-body");
export const messageInput = document.querySelector(".message-input");
export const sendMessageButton = document.querySelector("#send-message");

const socket = new WebSocket(`ws://${window.location.host}/ws`);

socket.onopen = () => {
  console.log("Websocket connection established.");
};

socket.onerror = (error) => {
  console.error("WebSocket Error:", error);
};

socket.onclose = (event) => {
  console.log("WebSocket Closed:", event.code, event.reason);
};

socket.onmessage = (event) => {
  try {
    const data = JSON.parse(event.data);
    console.log("Received message:", data);

    // Handle incoming messages
    if (data.message) {
      displayIncomingMessage(data.message);
    }
  } catch (error) {
    console.error("Error parsing incoming message:", error);
  }
};

const sendMessage = (message) => {
  if (socket.readyState === WebSocket.OPEN) {
    socket.send(JSON.stringify({ message })); // Only send message content
    console.log("Sent message:", message);
  } else {
    console.warn("WebSocket not open. ReadyState:", socket.readyState);
  }
};

const userData = {
  message: null,
};

// Create message element with dynamic classes and return it
export const createMessageElement = (content, ...classes) => {
  const div = document.createElement("div");
  div.classList.add("chat-message", ...classes);
  div.innerHTML = content;
  return div;
};

// Handle outgoing user messages
export const handleOutgoingMessage = (e) => {
  e.preventDefault();
  userData.message = messageInput.value.trim();
  messageInput.value = "";

  // Create and display user message
  const messageContent = `<div class="message-text"></div>`;
  const outgoingMessageDiv = createMessageElement(
    messageContent,
    "user-message"
  );
  outgoingMessageDiv.querySelector(".message-text").textContent =
    userData.message;
  chatBody.appendChild(outgoingMessageDiv);

  // Send the message via WebSocket
  sendMessage(userData.message);

  // Simulate a response with typing indicator after delay
  setTimeout(() => {
    // Create and display user message
    const messageContent = `<img class="bot-avatar" src="#" width="50" height="50" viewBox="0 0 1024 1024" />
          <div class="message-text">
            <div class="thinking-indicator">
              <div class="dot"></div>
              <div class="dot"></div>
              <div class="dot"></div>
            </div>
          </div>`;
    const incomingMessageDiv = createMessageElement(
      messageContent,
      "bot-message",
      "thinking"
    );
    chatBody.appendChild(incomingMessageDiv);
  }, 600);
};

export const displayIncomingMessage = (message) => {
  const messageContent = `<div class="message-text">${message}</div>`;
  const incomingMessageDiv = createMessageElement(
    messageContent,
    "bot-message"
  );
  chatBody.appendChild(incomingMessageDiv);
};

// Handle Enter key press for sending messages
messageInput.addEventListener("keydown", (e) => {
  const userMessage = e.target.value.trim();
  if (e.key === "Enter" && userMessage) {
    handleOutgoingMessage(e);
  }
});

sendMessageButton.addEventListener("click", (e) => handleOutgoingMessage(e));
