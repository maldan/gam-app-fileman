[...document.querySelectorAll("a[aria-label=фотография]")]
  .map((x) =>
    JSON.parse(
      x
        .getAttribute("onclick")
        .replace(/return showPhoto\('.*', '.*', (.*?), event\);?/gm, "$1")
    )
  )
  .map((x) => x.temp.w || x.temp.z || x.temp.y || x.temp.x)
  .join(" ");
