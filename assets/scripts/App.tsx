import React, { ReactElement } from 'react'
import { BrowserRouter, Switch, Route } from 'react-router-dom'
import PageHome from './pages/PageHome'

const App = (): ReactElement => {
  return (
    <BrowserRouter>
      <Switch>
        <Route path="/" exact component={PageHome} />
        <Route path="/login" exact component={PageHome} />
      </Switch>
    </BrowserRouter>
  )
}

export default App
