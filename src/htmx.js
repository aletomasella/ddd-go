import htmx from "htmx.org";

// Assinging htmx to window object
window.htmx = htmx;

const cookieToCheck = "refresh";

const cookie = document.cookie;

console.log(cookie);

let intervalId = null;

if (document.cookie.includes(cookieToCheck)) {
  intervalId = setInterval(() => {
    window.location.reload();
  }, 5000);
} else {
  if (intervalId) {
    clearInterval(intervalId);
  }
}
