export function UniversitySkeleton() {
  return (
    <div className="bg-white rounded-md shadow-md p-4 flex items-start animate-pulse">
      <div className="mr-4 bg-gray-300" style={{ width: 60, height: 60 }}></div>
      <div className="w-full">
        <div className="h-4 bg-gray-300 rounded w-3/4 mb-2"></div>
        <div className="h-3 bg-gray-300 rounded w-5/6"></div>
      </div>
    </div>
  );
}
