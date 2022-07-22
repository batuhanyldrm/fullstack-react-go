import React, {useState, useEffect} from 'react';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Blog from './Blog';
import axios from 'axios';



function InputField() {
    const [title, setTitle] = useState('')
    const [description, setDescription] = useState('')
    const [blogs, setBlogs] = useState('')
    
    const handleTitle = (e) => {
        setTitle(e.target.value)
        console.log(setTitle)
        console.log(title)
    }

    const handleDescription = (e) => {
        setDescription(e.target.value)
        console.log(setDescription)
        console.log(description)
    }

    const submit = () => {

        axios.post("http://localhost:3001/blogs", {
           title:title,
           description:description})
        .then(res => {
            console.log(res);
            console.log(res);
        })
        console.log(title,description)
    }
    const getBlogs = async () =>{
       const resp = await axios.get("http://localhost:3001/blogs")
       setBlogs(resp.data)
    }
    console.log(blogs)
    useEffect(() => {
      getBlogs();
      /* deleteRequestHandler(); */
    }, []);
    

        return(
            <div style={{display:"block", justifyContent:"center"}}>
                {blogs && blogs.map((blog) => <Blog blog ={blog} /* title={title} description={description} *//>)}
                    <TextField
                        id="title"
                        label="Title"
                        variant="outlined"
                        color="secondary"
                        value={title}
                        onChange={handleTitle}
                    />
                    <TextField
                        id="description"
                        label="Description"
                        variant="outlined"
                        color="secondary"
                        value={description}
                        onChange={handleDescription}
                    />
                    <Button variant="contained" color="primary" onClick={() => submit()}>submit</Button>
            </div>
        )
}

export default InputField;