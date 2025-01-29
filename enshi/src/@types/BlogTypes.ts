export type Blog = {
    blog_id: number;           
    user_id: number;        
    title: string;         
    description?: string;   
    category_id?: number;    
    created_at: Date;       
}