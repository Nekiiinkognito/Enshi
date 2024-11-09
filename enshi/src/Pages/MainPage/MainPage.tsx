import React from 'react'
import { Outlet } from 'react-router-dom'
import NavBar from '../../Components/NavBar/NavBar'
import { axiosLocalhost } from '../../api/axios/axios'

export default function MainPage() {
  return (
    <>
        <NavBar />
        <Outlet />
        <button
        onClick={
          async () => {
            let d = await axiosLocalhost.get("getCookie")
            console.log(d.data);
            
          }
        }>
          qwpofjqwifhqwuif
        </button>
    </>
  )
}
