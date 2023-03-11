import React from "react";
import { useNavigate } from "react-router-dom";
 
const CreateRoom = (props) =>{
    const navegate = useNavigate();
    const create = async (e) =>  {
        e.preventDefault()
        const resp = await fetch("http://localhost:8080/create"); 
        const {room_id} = await resp.json(); 
        navegate(`/room/${room_id}`)
    }
    return (
        <div>
            <button onClick={create}>Create</button>
        </div>
    )
}

export default CreateRoom