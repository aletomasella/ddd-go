import htmx from "htmx.org";

// Assinging htmx to window object
window.htmx = htmx;

(function () {
  let flag = false;

  async function checkToHotReload() {
    try {
      const response = await fetch("/reload");
      const data = await response.json();

      if (data.reload === "true") {
        window.location.reload();
      } else {
        flag = true;
      }
    } catch (error) {
      window.location.reload();
    }
  }

  const intervalId = setInterval(checkToHotReload, 1000);

  if (flag) {
    console.log("clearing interval");
    intervalId = clearInterval(intervalId);
  }
})();
