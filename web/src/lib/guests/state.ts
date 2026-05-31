export const loadStoredInvite = () => {
  return localStorage.getItem("invite") || "";
};

export const saveStoredInvite = (invite: string) => {
  localStorage.setItem("invite", invite);
};
