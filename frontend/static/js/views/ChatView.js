"use strict";

import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
    this.setTitle("Chat");
  }

  async getHtml() {
    return `
    <div class="chatbot-popup">
      <!-- Chat Header -->
      <div class="chat-header">
        <div class="header-info">
          <h2 style="color: #fff" class="logo-text">
            <i style="font-size: 20px" class="fas fa-comments chatbot-logo"></i>
            Forum Chat
          </h2>
        </div>
        <button id="close-chatbot" class="material-symbols-rounded">
          keyboard_arrow_down
        </button>
      </div>

      <!-- Chat Body -->
      <div class="chat-body">
        <div class="chat-message bot-message">
          <img
            class="bot-avatar"
            src="#"
            width="50"
            height="50"
            viewBox="0 0 1024 1024"
          />
          <div class="message-text">
            Hey there <br />
            How can I help you today?
          </div>
        </div>
      </div>

      <!-- Chat Footer -->
      <div class="chat-footer">
        <form action="#" class="chat-form">
          <textarea
            placeholder="Enter Message..."
            class="message-input"
            required
          ></textarea>
          <div class="chat-controls">
            <button type="button" class="material-symbols-rounded">
              sentiment_satisfied
            </button>
            <button type="button" class="material-symbols-rounded">
              attach_file
            </button>
            <button
              type="submit"
              id="send-message"
              class="material-symbols-rounded"
            >
              arrow_upward
            </button>
          </div>
        </form>
      </div>
    </div>
    `;
  }
}
