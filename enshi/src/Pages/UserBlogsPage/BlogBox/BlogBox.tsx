import { Card } from '@radix-ui/themes';
import { useNavigate } from 'react-router-dom';

type TBlogBox = {
    title?: string;
    blogId?: string;
}

export default function BlogBox(props: TBlogBox) {

    const navigate = useNavigate()

  return (
    <Card className='w-full h-20'
    onClick={() => navigate(``)}
    >
        {props?.title || "...No title..."}
        {props?.blogId || "adqwwd"}
    </Card>
  )
}
