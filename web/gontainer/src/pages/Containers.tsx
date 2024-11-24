import axios from 'axios';
import React, { useEffect, useState } from 'react'

interface Container {
  id: string;
  longId: string;
  name: string;
  image: string;
  status: string;
}

const Containers: React.FC = () => {

  const [containers, setContainers] = useState<Container[]>([]);
  const [isLoading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const retrieveContainerList = async () => {
      try {
        setLoading(true);
        const response = await axios.get<Container[]>('/gontainer/api/v1/containers');
        setContainers(response.data);
      } catch (error) {
        if (axios.isAxiosError(error)) {
          setError(error.message);
        } else {
          setError('Unexpected error!');
        }
      } finally {
        setLoading(false);
      }
    };

    retrieveContainerList();
  }, []);

  return (
    <div>
      Container!
      <ul>
        {containers.map( container => (
          <li key={container.id}>{container.name}</li>
        ))}
      </ul>
    </div>
  )
}

export default Containers