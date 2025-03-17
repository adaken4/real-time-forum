import AbstractView from "./AbstractView.js";

export default class SignupView extends AbstractView {
  constructor(params) {
    super(params);
    this.setTitle("Signup");
  }

  async getHtml() {
    return `
            <form action="/login" method="post" autocomplete="off" class="form">
              <p class="title">Register</p>
              <p class="message">Signup now and get full access to our app.</p>
              <div class="form-group">
                <label for="">
                  <input type="text" required />
                  <span>Firstname</span>
                </label>
                <label for="">
                  <input type="text" required />
                  <span>Lastname</span>
                </label>
              </div>
              <label for="">
                <input type="text" required />
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
              <p class="signin">Already have an account ? <a href="/login" data-link>Signin</a></p>
            </form>
        `;
  }

  async onMounted() {
    const form = document.getElementById("signupForm");
    form.addEventListener("submit", async (event) => {
      event.preventDefault();

      const userData = {
        firstName: document.getElementById("firstName").value,
        lastName: document.getElementById("lastName").value,
        email: document.getElementById("email").value,
        password: document.getElementById("password").value,
        passwordConfirm: document.getElementById("passwordConfirm").value,
      };

      try {
        const response = await fetch("/api/signup", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(userData),
        });

        if (response.ok) {
          alert("Signup successful!");
          window.location.href = "/login"; // Redirect to login
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
}
