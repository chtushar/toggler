import { queryKey } from "@/constants/queryKey";
import { ApiResponse } from "@/types";
import axios from "@/utils/axios";
import { useMutation, useQueryClient } from "@tanstack/react-query";

interface RegisterAdminData {
    name: string;
    email: string;
    password: string;
}

interface RegisterAdminResponseData {
    id: number;
    name: string;
    email: string;
    email_verified: boolean;
    role: string;
}

const useRegisterAdmin = () => {
    const client = useQueryClient();

    return useMutation({
        mutationFn: async (data: RegisterAdminData) => {
            const response = await axios.post("/api/add_admin", data)
            return response.data
        },
        onSuccess: async (data: ApiResponse<RegisterAdminResponseData>) => {
            if (data.success) {
                await client.setQueryData<RegisterAdminResponseData>(queryKey.user(), data.data)
            }
        }
    })
};

export default useRegisterAdmin;