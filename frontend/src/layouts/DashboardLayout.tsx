import { Outlet } from 'react-router-dom';
import Sidebar from '@/components/Sidebar';

const DashboardLayout = () => {
  return (
    <div className="flex">
      <Sidebar />
      <div className="flex-1 ml-0">
        <Outlet />
      </div>
    </div>
  );
};

export default DashboardLayout;
