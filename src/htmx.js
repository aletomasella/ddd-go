import htmx from "htmx.org";

// Assinging htmx to window object
window.htmx = htmx;

// async function checkToHotReload() {
//   try {
//     const response = await fetch("/reload");
//     const data = await response.json();

//     if (data.reload === "true") {
//       window.location.reload();
//     } else {
//       window.reloadIntervalId && clearInterval(window.reloadIntervalId);
//       window.reloadIntervalId = null;
//     }
//   } catch (error) {
//     window.location.reload();
//   }
// }

// window.reloadIntervalId = setInterval(checkToHotReload, 1000);
