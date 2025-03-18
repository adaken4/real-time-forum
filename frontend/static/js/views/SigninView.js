"use strict";

import AbstractView from "./AbstractView.js";

export default class SigninView extends AbstractView {
  constructor(params) {
    super(params);
    this.setTitle("Signin");
  }

  async getHtml() {
    return `
            <form id="signinForm" autocomplete="off" class="form">
              <p class="title">Signin</p>
              <p class="message">Access your account now.</p>
              <label for="">
                <input id="email" type="email" required />
                <span>Email</span>
              </label>
              <label for="">
                <input type="password" id="password" required />
                <span>Password</span>
                <span class="icon" id="togglePassword">
                  <i class="far fa-eye-slash"></i>
                </span>
              </label>
              <button class="submit">Signin</button>
              <p class="signin">Don't have an account? <a href="/signup" data-link>Signup</a></p>
            </form>
        `;
  }

  async onMounted() {
    const form = document.getElementById("signinForm");
    const passwordInput = document.getElementById("password");
    const togglePassword = document.getElementById("togglePassword");

    togglePassword.addEventListener("click", () => {
      this.togglePasswordVisibility(passwordInput, togglePassword);
    });

    form.addEventListener("submit", async (event) => {
      event.preventDefault();

      const userData = {
        email: document.getElementById("email").value,
        password: document.getElementById("password").value,
      };

      try {
        const response = await fetch("/api/signin", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(userData),
        });

        if (response.ok) {
          alert("Signin successful!");
          window.location.href = "/"; // Redirect to dashboard
        } else {
          const errorData = await response.json();
          alert(errorData.message || "Signin failed.");
        }
      } catch (error) {
        console.error("Error:", error);
        alert("Something went wrong.");
      }
    });
  }

  togglePasswordVisibility(inputElement, toggleElement) {
    if (inputElement.type === "password") {
      inputElement.type = "text";
      toggleElement.innerHTML = '<i class="far fa-eye"></i>';
    } else {
      inputElement.type = "password";
      toggleElement.innerHTML = '<i class="far fa-eye-slash"></i>';
    }
  }
}
