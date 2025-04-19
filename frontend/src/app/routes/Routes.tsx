import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { Error404Page } from '../../pages/Error404Page'
import { PagesListGrup } from '../../pages/PagesListGrup'
import { BlankPage } from '../../pages/BlankPage'

const AppRoutes = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<PagesListGrup />}>
          <Route index element={<PagesListGrup />} />
          <Route path="/list-grup" element={<PagesListGrup />} />
        </Route>
        <Route path="/blank" element={<BlankPage groupName="1ИСП9 - 45" />} />

        <Route path="*" element={<Error404Page />} />
      </Routes>
    </Router>
  )
}
export default AppRoutes
