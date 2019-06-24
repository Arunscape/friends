import React from 'react'
import ReactDOM from 'react-dom'
import App from 'app/App'
import * as serviceWorker from 'app/serviceWorker'

ReactDOM.render(<App />, document.getElementById('root'))

serviceWorker.unregister() // Change to register later to become a PWA
