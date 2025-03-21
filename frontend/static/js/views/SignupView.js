"use strict";

import AbstractView from "./AbstractView.js";
import { navigateTo } from "../index.js";

export default class SignupView extends AbstractView {
  constructor(params) {
    super(params);
    this.setTitle("Signup");
  }

  async getHtml() {
    return `
            <form id="signupForm" autocomplete="off" class="form form-styles">
              <p class="title">Register</p>
              <p class="message">Signup now and get full access to the Forum.</p>
              
              <label for="">
                <input id="nickname" type="text" required />
                <span>Nickname</span>
              </label>
              
              <div class="form-group">
                <label for="">
                  <input id="firstName" type="text" required />
                  <span>Firstname</span>
                </label>
                <label for="">
                  <input id="lastName" type="text" required />
                  <span>Lastname</span>
                </label>
              </div>

              <label for="">
                <input id="age" type="number" min="13" required />
                <span>Age</span>
              </label>

              <label for="">
                <select id="gender" required>
                  <option value="" disabled selected>Select Gender</option>
                  <option value="male">Male</option>
                  <option value="female">Female</option>
                  <option value="other">Other</option>
                </select>
                <span>Gender</span>
              </label>
              
              <label for="">
                <input id="email" type="text" required />
                <span>Email</span>
              </label>

              <label for="">
                <input type="password" id="password" required />
                <span>Password</span>
                <span class="icon" id="togglePassword">
                  <i class="far fa-eye-slash"></i>
                </span>
              </label>

              <label for="">
                <input type="password" id="passwordConfirm" required />
                <span>Confirm password</span>
                <span class="icon" id="togglePasswordConfirm">
                  <i class="far fa-eye-slash"></i>
                </span>
              </label>

              <button class="submit">Submit</button>
              <p class="signin">Already have an account ? <a href="/signin" data-link>Signin</a></p>
            </form>
        `;
  }

  async onMounted() {
    const form = document.getElementById("signupForm");

    const passwordInput = document.getElementById("password");
    const togglePassword = document.getElementById("togglePassword");
    const passwordConfirm = document.getElementById("passwordConfirm");
    const togglePasswordConfirm = document.getElementById(
      "togglePasswordConfirm"
    );

    togglePassword.addEventListener("click", () => {
      this.togglePasswordVisibility(passwordInput, togglePassword);
    });

    togglePasswordConfirm.addEventListener("click", () => {
      this.togglePasswordVisibility(passwordConfirm, togglePasswordConfirm);
    });

    form.addEventListener("submit", async (event) => {
      event.preventDefault();

      const userData = {
        nickname: document.getElementById("nickname").value,
        firstName: document.getElementById("firstName").value,
        lastName: document.getElementById("lastName").value,
        age: parseInt(document.getElementById("age").value, 10),
        gender: document.getElementById("gender").value,
        email: document.getElementById("email").value,
        password: document.getElementById("password").value,
        passwordConfirm: document.getElementById("passwordConfirm").value,
      };

      console.log(JSON.stringify(userData))

      try {
        const response = await fetch("/api/signup", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(userData),
        });

        if (response.ok) {
          console.log("Signup successful!");
          navigateTo("/signin"); // Redirect to signin
        } else {
          const errorData = await response.json();
          alert(errorData.message || "Signup failed.");
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
