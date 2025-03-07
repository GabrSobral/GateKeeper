import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useState,
} from "react";

interface SessionContextProps {
  session: GateKeeperSession | null;
  loading: boolean;
  error: GateKeeperSessionError | null;
  isAuthenticated: boolean;
}

interface SessionProviderProps {
  children: React.ReactNode;
}

const sessionContext = createContext<SessionContextProps>(
  {} as SessionContextProps
);

export function SessionProvider({ children }: SessionProviderProps) {
  const [session, setSession] = useState<GateKeeperSession | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<GateKeeperSessionError | null>(null);

  const fetchSession = useCallback(() => {
    setLoading(true);

    fetch("/api/session")
      .then(async (data) => {
        setSession(await data.json());
      })
      .catch((error) => {
        setError({ message: error.message });
      })
      .finally(() => {
        setLoading(false);
      });
  }, []);

  useEffect(() => fetchSession, [fetchSession]);

  return (
    <sessionContext.Provider
      value={{
        error,
        isAuthenticated: !!session,
        loading,
        session,
      }}
    >
      {children}
    </sessionContext.Provider>
  );
}

export const useSession = () => useContext(sessionContext);
