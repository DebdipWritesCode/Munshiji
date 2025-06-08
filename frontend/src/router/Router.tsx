import { Routes, Route, Navigate } from 'react-router-dom';
import Login from '../pages/Login';
import Signup from '../pages/Signup';
import SeeSheets from '../pages/SeeSheets';
import SheetDetails from '../pages/SheetDetails';
import NotFound from '../pages/NotFound';
import ProtectedRoute from './ProtectedRoute';

const Router = () => {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/login" />} />
      <Route path="/login" element={<Login />} />
      <Route path="/signup" element={<Signup />} />
      <Route
        path="/see-sheets"
        element={
          <ProtectedRoute>
            <SeeSheets />
          </ProtectedRoute>
        }
      />
      <Route
        path="/sheet/:id"
        element={
          <ProtectedRoute>
            <SheetDetails />
          </ProtectedRoute>
        }
      />
      <Route path="*" element={<NotFound />} />
    </Routes>
  );
};

export default Router;
