import CreatePostButton from "./CreatePostButton/CreatePostButton";
import UserButton from "./UserButton/UserButton";


export default function RightButtonBar() {
  return (
    <div className='flex flex-row justify-end flex-1 gap-4'>
        <CreatePostButton />
        <UserButton />
    </div>
  )
}