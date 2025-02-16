import { IApplication } from "@/services/dashboard/get-application-by-id";
import { ProviderCard } from "./ProviderCard";

type Props = {
  application: IApplication | null;
};

export function Providers({ application }: Props) {
  return (
    <div className="grid grid-cols-3 gap-3">
      {application?.oauthProviders.map((provider) => (
        <ProviderCard
          key={provider.id}
          title={provider.name}
          description={provider.description}
          isEnabled={provider.isEnabled}
          clientId={provider.clientId}
          clientSecret={provider.clientSecret}
        />
      ))}
    </div>
  );
}
