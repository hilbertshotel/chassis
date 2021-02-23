import os
import osproc
import strutils

import msg
import content

proc initProject() =
  if dirExists("frontend"):
    echo msg.existsError
    return
  
  let output = execProcess("tsc -v")
  if "Version" notin output:
    echo msg.warningTypescript
    return

  createDir("frontend/scripts")
  createDir("frontend/styles")
  writeFile("frontend/index.html", content.html)
  writeFile("frontend/styles/style.css", content.css)
  writeFile("frontend/scripts/script.ts", content.script)
  
  discard execCmd("tsc -init")
  echo msg.initialized  


proc main() =
  let args = commandLineParams()
  case args.len:
  of 0: initProject()
  else: echo msg.argumentsError


main()
