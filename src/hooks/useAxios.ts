import { useEffect, useState } from "react";
import axios, { AxiosResponse } from "axios";

export const useAxios = <T>() => {
  const [data, setData] = useState<T>();

  const fetch = (url: string, method: string): Promise<AxiosResponse<T>> => {
    return new Promise<AxiosResponse<T>>((resolve) => {
      axios
        .request({
          url,
          method,
        })
        .then((result) => {
          setData(result.data);
          resolve(result);
        });
    });
  };
  return { data, fetch };
};
