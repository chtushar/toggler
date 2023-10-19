import CodeMirror from '@uiw/react-codemirror'
import { javascript } from '@codemirror/lang-javascript'
import { githubLight } from '@uiw/codemirror-theme-github'

const js = `
    import React from 'react';
`

const Editor = () => {
  return (
    <CodeMirror
      value={js}
      height="200px"
      theme={githubLight}
      extensions={[javascript({ jsx: true })]}
    />
  )
}

export default Editor
