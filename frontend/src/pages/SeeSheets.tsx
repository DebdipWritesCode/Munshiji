import ToastComponent from "@/components/ToastComponent";
import { useEffect } from "react";

const fetchSheets = async () => {

}

useEffect(() => {
  fetchSheets();
}, [])

const SeeSheets = () => {
  return (
    <div className="min-h-screen bg-gradient-to-br from-[#fbfcff] via-[#EDEFFF] to-[#D4E0FF]">


      <ToastComponent />
    </div>
  );
};

export default SeeSheets;
