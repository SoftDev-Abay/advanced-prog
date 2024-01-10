import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

const Movie = () =>{
    const[movie, setMovie] = useState({});
    let{ id } = useParams();

    useEffect(() => {
       let myMovie = {
            id:1,
            title : "Highlender",
            release_date: "1986-03-06",
            runtime: 166,
            mpaa_rating: "R",
            description: "SOmething here",
        } 
        setMovie(myMovie)
    }, [id])
    return (
            <div>
                <h2>Movie: {movie.title}</h2>
                <small><em> {movie.release_date}, {movie.runtime} minuts, Rated - {movie.mpaa_rating}</em></small>
                <hr />
                <p> {movie.description}</p>
            </div>
    )
}

export default Movie;