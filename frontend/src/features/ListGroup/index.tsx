import { Group } from "@/entities/group/types";
import { GroupCard } from "@/widgets/GroupCard";
import { Pagination } from "@/shared/ui/Pagination";
import { useEffect, useMemo, useState } from "react";

const PER_PAGE = 8;

interface ListGroupProps {
  groups: Group[];
}

export function ListGroup({ groups }: ListGroupProps) {
  const [page, setPage] = useState(1);

  useEffect(() => setPage(1), [groups]);

  const paginated = useMemo(() => {
    const start = (page - 1) * PER_PAGE;
    return groups.slice(start, start + PER_PAGE);
  }, [groups, page]);

  return (
    <>
      {paginated.length > 0 &&
        paginated.map((group: Group) => (
          <GroupCard key={group.id} group={group} />
        ))}
      <div className="mt-2 flex items-center justify-center">
        <Pagination
          page={page}
          pages={Math.ceil(groups.length / PER_PAGE)}
          onChange={setPage}
        />
      </div>
    </>
  );
}
