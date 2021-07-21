import React, { ReactElement } from 'react'
import { BrowserRouter, Switch, Route } from 'react-router-dom'
import PageLogin from './pages/Login'
import PageRegister from './pages/Register'
import PageHome from './pages/PageHome'

const App = (): ReactElement => {
  return (
    <div className="wrapper min-h-full flex flex-col">
      <nav className="flex-none shadow-lg border-b border-black border-opacity-10 z-30">
        <div className="container mx-auto px-4 md:px-0 flex">
          {/* Branding */}
          <div className="block flex-1 h-16 py-2">
            <a href="/" className="inline-block h-full w-auto">
              <img src="/assets/images/logo.svg" alt="LoginHub" className="block h-full" />
            </a>
          </div>

          {/* Profile Icon */}
          <a href="/profile" className="block w-16 h-16 p-4 text-center hover:bg-opacity-20 hover:bg-black transition-colors">
            <img
              src="/assets/images/outline_account_circle_black_48dp.png"
              className="block h-8 w-8 rounded-full"
              alt="Guest" />
          </a>
        </div>
      </nav>

      {/* Main Contents */}
      <BrowserRouter>
        <Switch>
          <Route path="/" exact component={PageHome} />
          <Route path="/login" component={PageLogin} />
          <Route path="/register" component={PageRegister} />
        </Switch>
      </BrowserRouter>
    </div>
  )
}

export default App
