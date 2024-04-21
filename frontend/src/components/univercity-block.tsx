
import { Univercity } from "@/api/university";
import { API_HOST } from "@/lib/auth";
import  Image  from "next/image";

interface UniversityBlockParams {
  universityId: string;
}

async function getUniversity(universityId: string) {
  const response = await fetch(
    `${API_HOST}/api/v1/universities/${universityId}`,
    { cache: "no-cache" }
  );
  return response.json();
}

export default async function UniversityBlock({
  universityId,
}: UniversityBlockParams) {
  const response  = await getUniversity(universityId);
  const univercity: Univercity = response.data
  console.log(univercity);
  return (
    <div className="flex flex-col w-2/3 items-center bg-slate-100 p-20 m-20 rounded-xl">
      <img src={univercity.logo} width={0} height={0} style={{width:'120px', height: "120px"}} className="rounded-full bg-slate-300 m-4 "alt="logo"/>
      <p className="text-4xl font-bold">{univercity.name}</p>
      <p className="">{univercity.longName}</p>
      <p className="">{univercity.region}</p>
      <p className="">{univercity.type}</p>
      <p className="text-5xl font-bold m-2">{univercity.rating}</p>
      <p>Рейтинг</p>
      <p className="text-5xl font-bold m-2">{univercity.studyFields}</p>
      <p>Направлений</p>
      <p className="text-5xl font-bold m-2">{univercity.budgetPlaces}</p>
      <p>Бюджетные места</p>
    </div>
  );
}
