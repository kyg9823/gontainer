import React from 'react'
import { Link, Router } from 'react-router-dom'

const Top: React.FC = () => {
  return (
    <nav className="top">
      <ul className="top-menu">
        <li>
          <Link to="/containers">Containers</Link>
        </li>
        <li>
          <Link to="/images">Images</Link>
        </li>
      </ul>
    </nav>
  )
}

export default Top