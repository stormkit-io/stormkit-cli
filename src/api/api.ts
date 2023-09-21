const baseUrl = "https://api.stormkit.io/v1";

const headers = () => ({
  Authorization: `${process.env.SK_API_KEY}`,
});

export default {
  get<ReturnType>(url: string, options?: RequestInit) {
    return fetch(`${baseUrl}${url}`, {
      ...options,
      headers: headers(),
      method: "GET",
    }).then(async (res) => {
      return (await res.json()) as ReturnType;
    });
  },
};
