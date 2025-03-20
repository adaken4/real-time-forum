"use strict";
import Dashboard from "./views/Dashboard.js";
import Posts from "./views/Posts.js";
import PostView from "./views/PostView.js";
import Settings from "./views/Settings.js";
import SignupView from "./views/SignupView.js";
import SigninView from "./views/SigninView.js";

const pathToRegex = (path) =>
  new RegExp("^" + path.replace(/\//g, "\\/").replace(/:\w+/g, "(.+)") + "$");

const getParams = (match) => {
  const values = match.result.slice(1);
  const keys = Array.from(match.route.path.matchAll(/:(\w+)/g)).map(
    (result) => result[1]
  );

  return Object.fromEntries(
    keys.map((key, i) => {
      return [key, values[i]];
    })
  );
};

export const navigateTo = (url) => {
  history.pushState(null, null, url);
  router();
};

const isAuthenticated = () => {
  console.log(document.cookie.includes("session_id"));
  return document.cookie.includes("session_id"); // Simple check
};

const router = async () => {
  const routes = [
    { path: "/", view: isAuthenticated() ? Dashboard : SigninView },
    { path: "/signup", view: SignupView },
    { path: "/signin", view: SigninView },
    { path: "/posts", view: Posts },
    { path: "/posts/:id", view: PostView },
    { path: "/settings", view: Settings },
  ];

  // Test each route for potential match
  const potentialMatches = routes.map((route) => {
    return {
      route: route,
      result: location.pathname.match(pathToRegex(route.path)),
    };
  });

  let match = potentialMatches.find(
    (potentialMatch) => potentialMatch.result !== null
  );
  if (!match) {
    match = {
      route: routes[0],
      result: location.pathname,
    };
  }

  const view = new match.route.view(getParams(match));

  document.querySelector("#app").innerHTML = await view.getHtml();
  if (typeof view.onMounted === "function") {
    view.onMounted(); // Call onMounted() if defined
  }
};

window.addEventListener("popstate", router);

document.addEventListener("DOMContentLoaded", () => {
  document.body.addEventListener("click", (e) => {
    if (e.target.matches("[data-link]")) {
      e.preventDefault();
      navigateTo(e.target.href);
    }
  });
  router();
});

async function checkAuthStatus() {
  try {
    const response = await fetch("/api/auth/status", {
      credentials: "include", // Send cookies
    });
    const data = await response.json();

    if (data.authenticated) {
      console.log("User is logged in:", data.user_id);
      return data.user_id;
    } else {
      console.log("User is not authenticated.");
      return null;
    }
  } catch (error) {
    console.error("Error checking auth status:", error);
    return null;
  }
}

async function protectRoute(requiredAuth, redirectTo = "/signin") {
  const userID = await checkAuthStatus();

  if (requiredAuth && !userID) {
    console.log("Redirecting to signin...");
    navigateTo(redirectTo);
  }
}

const socket = new WebSocket(`ws://${window.location.host}/ws`);

socket.onopen = (event) => {
  console.log("Websocket connection established");
  const message = {"message": "Hello, fellow user!"};
  sendMessage(message);
}

socket.onmessage = (event) => {
  try {
    const data = JSON.parse(event.data);
    console.log("Received message:", data);
  } catch (error) {
    console.error("Error parsing incoming message:", error);
  }
};

function sendMessage(message) {
  if (socket.readyState === WebSocket.OPEN) {
    socket.send(JSON.stringify(message)); // Only send message content
    console.log("Sent message:", message)
  } else {
    console.warn("WebSocket not open. ReadyState:", socket.readyState);
  }
};