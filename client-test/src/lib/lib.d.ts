type GateKeeperSession = {
  user: {
    id: string;
    firstName: string;
    lastName: string;
    email: string;
    displayName: string;
    applicationId: string;
  };
  accessToken: string;
};

type GateKeeperSessionError = {
  message: string;
};
