import { AvatarImage, AvatarFallback, Avatar } from "@/components/ui/avatar";
import { ReplyData } from "@/lib/use-review";
import { Skeleton } from "./ui/skeleton";

// Interface for the reply data
interface ReplyProps {
  reply: ReplyData;
}

export function ReplySkeleton() {
  return (
    <div className="max-w-4xl mx-auto p-4 rounded-lg shadow-md w-full">
      <div className="flex items-start space-x-4">
        <Skeleton className="w-12 h-12" />
        <div className="flex-1">
          <Skeleton className="h-6 mb-2 w-1/4" />
          <Skeleton className="h-6 w-3/4" />
        </div>
      </div>
    </div>
  );
}

// The Reply component for rendering a reply to a review
export function Reply({ reply }: ReplyProps) {
  return (
    <div className="max-w-4xl mx-auto p-4 rounded-lg shadow-md w-full">
      <div className="flex items-start space-x-4">
        <Avatar>
          <AvatarImage
            alt="User Avatar"
            src="/placeholder.svg?height=50&width=50"
          />
          <AvatarFallback>{reply.name.charAt(0)}</AvatarFallback>
        </Avatar>
        <div className="flex-1">
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-2">
              <span className="text-sm text-muted-foreground">
                {reply.date}
              </span>
            </div>
          </div>
          <div className="mt-4">
            <p className="mt-2">{reply.text}</p>
          </div>
        </div>
      </div>
    </div>
  );
}
